# E2Easy-Go

**E2Easy-Go** is a library to create a complex end-to-end tests and validations writing a simple yaml file.
It's possible to define a set of information based on pre-defined configuration to start a flow of REST API calls, and use the result
for validations or next flows call.

The library runs on the console where it is possible to pass an input file (.yaml) 
to be consumed, this file defines the steps and validations of the end-to-end test.

###### How to execute:
![execution](https://github.com/carloshjoaquim/E2Easy-Go/blob/master/images/cmd.png?raw=true)


### Allowed properties
_____

| Property   | Information                                      | Example                                                                                                            |
|------------|--------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| name       | Set the E2E flow name                            | name: my_e2e_test                                                                                                  |
| steps      | Array of steps for E2E Test                      | steps:<br>  - name: ....                                                                                           |
| steps.name | Name of a single E2E Test step                   | steps:<br>  - name: get_validation                                                                                 |
| steps.path | URL path for call API                            | steps:<br>  - ...<br>    path: http://my-endpoint/valid                                                            |
| method     | REST Method to call API for this step            | steps:<br>  - ...<br>    method: GET                                                                               |
| body       | string with json BODY to cal API if necessary    | steps:<br>  - ...<br>    body: "{ \"reference_id\": \"TST-E2E004\"}"                                               |
| headers    | header information if necessary (name and value) | steps:<br>  - ...<br>    headers:<br>      - name: "Authorization<br>        value: Z8LM1y7LMTgbUFgJ4maMKGTdsIR8Nb |

#### Example 1:
##### To create a simple GET call to an URL and validate the result:
 
````
name: contract_retention
steps:
  - name: contrato
    path: https://mail-generator.herokuapp.com/generate?domain=hotmail
    method: GET
    vars:
      fullBody: response.body
      email: response.body.mail
      statusCode: response.statusCode
    tests:
      - name: return_hotmal
        expected: "hotmail"
        actual: ${email}
        type: contains
      - name: status_ok
        expected: 200
        actual: ${statusCode}
        type: equals
``