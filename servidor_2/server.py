
#import mysql.connector
#python -m grpc_tools.protoc -I="servidor_2" --python_out=servidor_2 --grpc_python_out=servidor_2  "servidor_2/servicios.proto"
import MySQLdb
import servicios_pb2
import servicios_pb2_grpc 
import grpc
from concurrent import futures

myDB = MySQLdb.connect( host='localhost', user= 'admin', passwd='12345678', db='db_inventario' )
cur = myDB.cursor()


class ProveedorServicer(servicios_pb2_grpc.ProveedorServicer):
    def SuministrarProductos(self, request, context):
      print(request)
      respuesta = servicios_pb2.Respuesta()
      respuesta.ack = f"{request.id}"
      id = str(request.id)
      sql = "SELECT * FROM inventario WHERE id_producto = %s"
      cur.execute(sql, id)
      req = cur.fetchall()
      producto = req[0]
      if producto[2] < request.cantidad:
        CANTIDAD = 5
      else:
        CANTIDAD = producto[2]-request.cantidad
      sql = "UPDATE inventario SET cantidad_disponible = %s WHERE id_producto = %s"
      cur.execute(sql,(CANTIDAD,request.id))
      myDB.commit()
      print("Se suministraron ",request.cantidad," productos.")
      return respuesta

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
  print("Server inicializado.")
  servicios_pb2_grpc.add_ProveedorServicer_to_server(ProveedorServicer(),server)
  server.add_insecure_port('0.0.0.0:9000')
  server.start()
  server.wait_for_termination()

serve()
myDB.close()