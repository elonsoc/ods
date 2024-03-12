# migrate.py is a

import pymssql as mssql
import psycopg2 as pg


def main():
    mssql_conn = mssql.connect("localhost", "sa", "Password123", "data")
    pg_conn = pg.connect("dbname=postgres user=postgres password=postgrespw host=localhost port=5432")

    mssql_cur = mssql_conn.cursor()


    a = mssql_cur.execute("SELECT * FROM dbo.buildings")

    b = a.fetchall()

    print(b)
