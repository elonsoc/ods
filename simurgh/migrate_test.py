# TODO: periodic update

import pandas as pd
from sqlalchemy import create_engine, exc, create_engine, exc, MetaData, Table, Column, VARCHAR, Numeric, TIMESTAMP, text

# define connection variables
mssql_user = "sa"
mssql_pass = "Password123"
mssql_host = "localhost:1433"
mssql_database = "data"

psql_user = "postgres"
psql_pass = "postgrespw"
psql_host = "localhost:5432"
psql_database = "postgres"

def main():
    try:
        mssql_engine = connect_mssql()
        print("Connected to MSSQL successfully.")

        psql_engine, buildings_table_psql = connect_psql()
        print("Connected to PostgreSQL successfully.")

        with psql_engine.connect() as psql_conn:
                    # create PSQL buildings table, if doesn't exist
                    buildings_table_psql.create(psql_engine)  
                    # fetch mssql data
                    mssql_data = pd.read_sql("SELECT * FROM dbo.buildings", mssql_engine)
                    # insert into PSQL buildings table
                    mssql_data.to_sql('buildings', psql_conn, index=False, if_exists='replace', method='multi')

        print("Data copied from MSSQL to PostgreSQL.")

    except exc.OperationalError as e:
        print(f"Connection failed: {e}")
        

def connect_mssql():
    # define mssql connection object and engine
    mssql_url_object = f"mssql+pyodbc://{mssql_user}:{mssql_pass}@{mssql_host}/{mssql_database}?driver=ODBC+Driver+17+for+SQL+Server"
    mssql_engine = create_engine(mssql_url_object)
    mssql_engine.connect()

    return mssql_engine


def connect_psql():
    # define psql connection object and engine
    psql_url_object = f"postgresql://{psql_user}:{psql_pass}@{psql_host}/{psql_database}"
    psql_engine = create_engine(psql_url_object)
    # define PSQL table structure
    buildings_table_psql = Table(
        'buildings', MetaData(),
        Column('BUILDINGS_ID2', VARCHAR(4), primary_key=True, nullable=False),
        Column('BLDG_DESC', VARCHAR(50)),
        Column('BLDG_LOCATION', VARCHAR(5)),
        Column('BLDG_LOCATION_REPRESENTATION', VARCHAR(30)),
        Column('BLDG_TYPE', VARCHAR(10)),
        Column('BLDG_TYPE_REPRESENTATION', VARCHAR(32)),
        Column('BLDG_LONG_DESC', VARCHAR(1996)),
        Column('BLDG_CITY', VARCHAR(25)),
        Column('BLDG_STATE', VARCHAR(2)),
        Column('BLDG_ZIP', VARCHAR(10)),
        Column('BLDG_SECTOR', VARCHAR(10)),
        Column('BLDG_SECTOR_REPRESENTATION', VARCHAR(32)),
        Column('BLDG_LATITUDE', Numeric),
        Column('BLDG_LONGITUDE', Numeric),
        Column('BUILDINGS_ADD_DATE', TIMESTAMP),
        Column('BUILDINGS_CHGDATE', TIMESTAMP)
    )
    return psql_engine, buildings_table_psql

if __name__ == '__main__':
     main()

