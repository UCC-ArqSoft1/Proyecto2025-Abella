create database gymdb;
use gymdb;

// Una vez creada la base se datos y hecho el migrate se puede continuar con las consultas:

INSERT INTO activity_types (name) VALUES ("Acuaticas");

INSERT INTO user_types (type) values("Usuario");
INSERT INTO user_types (type) values("Admin");

INSERT INTO users (user_type_id,email,hashed_password,name,last_name,documentation,is_coach) values(2,"root@gym.com","4813494d137e1631bba301d5acab6e7bb7aa74ce1185d456565ef51d737677b2","root","root",1,1); // Usuario tipo admin que tiene todos los permisos.

INSERT into activities (name,description,activity_type_id,duration,capacity,coach_id) values("Natacion","Actividad acuantica en nuestras piscinas locales!.",1,2,20,1); // Actividad generica con fallo de ortografia en acuatica para probar lo del tipo de actividad

// Por ahora el resto de consultas que se hacen estan implementadas en el backend y se pueden acceder mediante el front.
// Consultar este archivo una ves este implementada la funcionalidad de inscripciones ya que habra que hacer una manual para probar el sistema.


// Las consultas presentes en este archivo no son suficientes para demostrar la capacidad total de la app son solo para poder arrancar a utilizarla. una ves la app este corriendo lo ingresado aqui se puede borrar. 