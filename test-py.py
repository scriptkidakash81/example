import mysql.connector
import logging
import os

# Hardcoded credentials (Vulnerability 1)
db_config = {
    'user': 'admin',
    'password': 'mysecretpassword',
    'host': 'localhost',
    'database': 'mydb'
}

# Potential SQL injection vulnerability (Vulnerability 2)
user_input = "Robert'); DROP TABLE users; --"
query = f"SELECT * FROM users WHERE name = '{user_input}'"

try:
    # Insecure database connection (Vulnerability 3)
    cnx = mysql.connector.connect(**db_config)
    cursor = cnx.cursor()

    # Insecure logging of sensitive data (Vulnerability 4)
    logging.warning(f"Connecting to database with credentials: {db_config}")

    # Potential credentials exposure through environment variables (Vulnerability 5)
    print(f"DB_PASSWORD: {os.environ.get('DB_PASSWORD')}")

    # Execute the query
    cursor.execute(query)

    # Fetch and print the results
    for row in cursor.fetchall():
        print(row)

except mysql.connector.Error as err:
    print(f"Error: {err}")

finally:
    cursor.close()
    cnx.close()
