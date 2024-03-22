# App logic:

1. React as frontend
2. Golang as backend

## APIs

| n 	| Name              	| Arguments                    	| Return                     	| URL                       	|
|---	|-------------------	|------------------------------	|----------------------------	|---------------------------	|
| 1 	| Login             	| username, password           	| auth_temporary_key         	| /apiv1/login/             	|
| 2 	| User data         	| username, auth_temporary_key 	| email, name, surname, role 	| /apiv1/user-data/         	|
| 3 	| Own student marks 	| username, auth_temporary_key 	| [] marks                   	| /apiv1/own-student-marks/ 	|
