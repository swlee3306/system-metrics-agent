"""Checks public configuration without contacting an operating host."""
from pathlib import Path
import subprocess
import unittest

ROOT = Path(__file__).resolve().parents[1]

class PublicConfigTests(unittest.TestCase):
    def test_scripts_require_explicit_target(self):
        for name in ('service-active.sh', 'service-deactive.sh'):
            script = ROOT / 'script' / 'arm64' / name
            source = script.read_text()
            # Refuse to execute a legacy script that could contact a default host.
            self.assertTrue('SERVER_IP=${2:?Specify the target host}' in source, 'missing host guard')
            self.assertTrue('SSH_KEY_PATH=${4:?Specify the SSH key path}' in source, 'missing key guard')
            self.assertNotIn('/Users/', source)
            result = subprocess.run(['bash', str(script)], capture_output=True, timeout=5)
            self.assertNotEqual(result.returncode, 0)

    def test_datasource_passwords_are_placeholders(self):
        lines = [line.strip() for line in (ROOT / 'grafana/datasource.yaml').read_text().splitlines()
                 if line.strip().startswith('password:')]
        self.assertEqual(len(lines), 2)
        self.assertTrue(all('REPLACE_WITH_' in line for line in lines))

    def test_passwords_use_secure_json_data(self):
        lines = (ROOT / 'grafana/datasource.yaml').read_text().splitlines()
        passwords = [i for i, line in enumerate(lines) if line.strip().startswith('password:')]
        self.assertEqual(len(passwords), 2)
        for i in passwords:
            self.assertEqual(lines[i - 1], '    secureJsonData:')
            self.assertTrue(lines[i].startswith('      password:'))

if __name__ == '__main__':
    unittest.main()
