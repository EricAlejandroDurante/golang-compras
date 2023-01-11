import mysql.connector
import time
import random as r

myDB = mysql.connector.connect(host='localhost', user= 'admin', passwd='12345678', db='db_inventario')
cur = myDB.cursor()

while True:
    consulta = cur.execute("SELECT * FROM inventario WHERE cantidad_disponible=0")
    myresult = cur.fetchall()
    if(len(myresult)!=0):
        sql = "UPDATE inventario SET cantidad_disponible = %s WHERE id_producto=%s"
        for i in myresult:
            id = i[0]
            cantidad = r.randint(1,10)
            val = (cantidad,id)
            req = cur.execute(sql,val)
            myDB.commit()
    time.sleep(60)