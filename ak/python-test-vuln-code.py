Python:

import os
import subprocess
import sqlite3
import hashlib

# Vulnerability 1: Hardcoded credentials
username = "admin"
password = "mysecretpassword"

# Vulnerability 2: Insecure password hashing
hashed_password = hashlib.md5(password.encode()).hexdigest()

# Vulnerability 3: SQL injection
conn = sqlite3.connect("example.db")
cursor = conn.cursor()
user_input = "Robert'); DROP TABLE users; --"
query = f"SELECT * FROM users WHERE name = '{user_input}'"
cursor.execute(query)

# Vulnerability 4: Command injection
subprocess.run(f"echo {user_input}", shell=True)

# Vulnerability 5: Insecure file permissions
with open("sensitive_data.txt", "w") as f:
    f.write("Top secret data")

# Vulnerability 6: Insecure temporary file creation
import tempfile
tmp_file = tempfile.NamedTemporaryFile()
tmp_file.write(b"sensitive data")
tmp_file.close()

# Vulnerability 7: Potential path traversal vulnerability
file_path = "../sensitive_data.txt"
with open(file_path, "r") as f:
    print(f.read())

# Vulnerability 8: Insecure use of eval()
eval(user_input)

# Vulnerability 9: Potential cross-site scripting (XSS) vulnerability
print(f"<script>alert('{user_input}')</script>")

# Vulnerability 10: Insecure use of pickle
import pickle
data = pickle.loads(b'x\x01\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00')

# Additional vulnerabilities
# Vulnerability 11: Insecure SSL/TLS configuration
import ssl
ssl_context = ssl.create_default_context()
ssl_context.check_hostname = False
ssl_context.verify_mode = ssl.CERT_NONE
