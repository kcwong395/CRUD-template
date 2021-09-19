import psycopg2

# init the connection
conn = psycopg2.connect(
    host = "ubuntu-VirtualBox",
    database = "testDb",
    user = "postgres",
    password = "postgres",
    port = "5000"
)

cur = conn.cursor()

cur.execute("create table employees (id integer not null, name varchar(20) not null, primary key (id))")

cur.execute("insert into employees (id, name) values (%s, %s)", (1, 'Bob'))

cur.execute("select id, name from employees")

rows = cur.fetchall()
for r in rows:
    print(f"id {r[0]} name {r[1]}")

# commit changes
conn.commit()

cur.close()

# close the connection
conn.close()