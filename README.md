# SEMANTIC VERSIONING - SEMVER

https://en.wikipedia.org/wiki/Software_release_life_cycle
https://semver.org/
https://en.wikipedia.org/wiki/Software_versioning
https://jsonresume.org/schema/


https://github.com/shurcooL/gostatus
https://github.com/pypa/pip/blob/master/pyproject.toml
https://github.com/russross/blackfriday

major = incompatible
minor = aditions / changes 
patch = bugs fixes 

1.- crear un svg de las ultima version lanzada 
2.- X.Y.Z - todos deben ser enteros positivos, incrementos de 1 
    x = major , y = minor & z =path  
3.- cuando se lanza un relanse no se modifica se debe poner un tag. 
4.- X =0 pre-alfa o alfa no estable 
5.- x=1 rc-beta
6.- Z se incrementa cuando son compatibles con la version anterior 
7.- Y se incrementa cuando se introduce una nueva funcionalidad  compatible con la v anterior
8.- X se incrementa cuando no es compatible con la version anterior. Los ZY se reinician en cero 
9.- Pre-release se representa por - despues del patch  formato [0-9A-Za-z-]
10.-build se represneta por + depues del path o pre-release no hace diferencia de versionado 
11.-precedencias 
    
    pre-alfa = se refiere a todas las actividades realizadas durante el proyecto de software antes de las pruebas formales.
    
    alfa = En esta fase, los desarrolladores generalmente prueban el software
    
    beta =  comienza cuando la función del software está completa, pero es probable que contenga una serie de errores conocidos o desconocidos.
    
    Perpetual beta = se agregan continuamente nuevas funciones al software sin establecer una versión final "estable"

    Open and closed beta = Las versiones beta cerradas se lanzan a un grupo restringido de personas para una prueba de usuario por invitación, mientras que los probadores beta abiertos pertenecen a un grupo más grande o cualquier persona interesada.

    Release candidate (silver) =  es una versión beta con potencial para ser un producto estable, que está listo para su lanzamiento a menos que surjan errores importantes.

    Stable release =  versiones normales o estables 
                        versiones de soporte a largo plazo (LTS) 

    discontinued =Cuando el software ya no se vende ni se admite, 