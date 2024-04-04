from dotenv import load_dotenv
import os

import pandas as pd
from sqlalchemy import create_engine, exc, create_engine, exc, MetaData, Table, Column, VARCHAR, Numeric, TIMESTAMP, text
        
load_dotenv()

mssql_user = os.environ['MSSQL_USER']
mssql_pass = os.environ['MSSQL_PASSWORD']
mssql_host = os.environ['MSSQL_HOST']
mssql_database = os.environ['MSSQL_DATABASE']

def connect_mssql():
    # define mssql connection object and engine
    # https://learn.microsoft.com/en-us/sql/connect/odbc/download-odbc-driver-for-sql-server?view=sql-server-ver16
    mssql_url_object = f"mssql+pyodbc://{mssql_user}:{mssql_pass}@{mssql_host}/{mssql_database}?driver=ODBC+Driver+17+for+SQL+Server"
    mssql_engine = create_engine(mssql_url_object)
    mssql_engine.connect()

    return mssql_engine

def connect_psql():
    '''these are the psql creds'''
    psql_user = "postgres"
    psql_pass = "postgrespw"
    psql_host = "localhost:5432"
    psql_database = "postgres"

    # define psql connection object and engine
    psql_url_object = f"postgresql://{psql_user}:{psql_pass}@{psql_host}/{psql_database}"
    psql_engine = create_engine(psql_url_object)
    # define PSQL table structure

    # buildings_table_psql = """
    # CREATE TABLE IF NOT EXISTS buildings (
    #     buildings_id VARCHAR(4) PRIMARY KEY NOT NULL,
    #     bldg_desc VARCHAR(50),
    #     bldg_location VARCHAR(5),
    #     bldg_location_representation VARCHAR(30),
    #     bldg_type VARCHAR(10),
    #     bldg_type_representation VARCHAR(32),
    #     bldg_long_desc VARCHAR(1996),
    #     bldg_city VARCHAR(25),
    #     bldg_state VARCHAR(2),
    #     bldg_zip VARCHAR(10),
    #     bldg_sector VARCHAR(10),
    #     bldg_sector_representation VARCHAR(32),
    #     bldg_latitude NUMERIC,
    #     bldg_longitude NUMERIC,
    #     buildings_add_date TIMESTAMP,
    #     buildings_chgdate TIMESTAMP
    # );

    buildings_table_psql = Table(
        'buildings', MetaData(),
        Column('buildings_id', VARCHAR(4), primary_key=True, nullable=False),
        Column('bldg_desc', VARCHAR(50)),
        Column('bldg_location', VARCHAR(5)),
        Column('bldg_location_representation', VARCHAR(30)),
        Column('bldg_type', VARCHAR(10)),
        Column('bldg_type_representation', VARCHAR(32)),
        Column('bldg_long_desc', VARCHAR(1996)),
        Column('bldg_city', VARCHAR(25)),
        Column('bldg_state', VARCHAR(2)),
        Column('bldg_zip', VARCHAR(10)),
        Column('bldg_sector', VARCHAR(10)),
        Column('bldg_sector_representation', VARCHAR(32)),
        Column('bldg_latitude', Numeric),
        Column('bldg_longitude', Numeric),
        Column('buildings_add_date', TIMESTAMP),
        Column('buildings_chgdate', TIMESTAMP)
    )

    return psql_engine, buildings_table_psql

def main():

    try:
        mssql_engine = connect_mssql()
        print("Connected to MSSQL successfully.")

        psql_engine, buildings_table_psql = connect_psql()
        print("Connected to PostgreSQL successfully.")
        
        with psql_engine.connect() as psql_conn: 
            # fetch mssql data
            mssql_data = pd.read_sql("SELECT * FROM dbo.vw_ESC_BUILDINGS", mssql_engine)
            # insert into PSQL buildings table -- make if it doesn't exist, replace/overwrite table if it does
            mssql_data.to_sql('buildings', psql_conn, index=False, if_exists='replace', method='multi')

        print("Data copied from MSSQL to PostgreSQL.")

    except exc.OperationalError as e:
        print(f"Connection failed: {e}")

if __name__ == '__main__':
     main()