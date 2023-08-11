import psycopg2
# DB_HOST=localhost
# DB_NAME=fiber-app-site-for-pasha
# DB_USER=postgres
# DB_PASSWORD=password
# DB_PORT=5432
class db():
    def __init__(self):
        self.dbname = 'fiber-app-site-for-pasha'
        self.user = 'postgres'
        self.password = 'password'
        self.host = 'localhost'
        self.table_name='notes'
    def connect(self):
        conn = psycopg2.connect(dbname=self.dbname, user=self.user,
                                password=self.password, host=self.host)
        # self.cursor = self.conn.cursor()
        return conn
    def select_all(self):
        cursor = self.connect().cursor()
        cursor.execute('SELECT * FROM notes')
        records = cursor.fetchall()
        return records
    def delet_all(self):
        conn = self.connect()
        cursor = conn.cursor()
        cursor.execute(f"""DELETE * FROM {self.table_name}""")
        cursor.execute('SELECT * FROM notes LIMIT 10')
        conn.commit()
        cursor.close()
        conn.close()
    def del_row(self, id_key):
        conn = self.connect()
        cursor = conn.cursor()
        cursor.execute(f"""DELETE FROM {self.table_name} WHERE id = '{id_key}'""")
        cursor.execute('SELECT * FROM notes LIMIT 10')
        conn.commit()
        cursor.close()
        conn.close()


if __name__ == '__main__':
    a=db()
    for i in a.select_all():
        print(i)
    print()
    print(a.select_all()[0][0])
    a.del_row(a.select_all()[0][0])
    print()
    for i in a.select_all():
        print(i[0])