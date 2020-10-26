# E2Easy-Go

**E2Easy-Go** is a library to create a complex end-to-end tests and validations writing a simple yaml file.
It's possible to define a set of information based on pre-defined configuration to start a flow of REST API calls, and use the result
for validations or next flows call.

The library runs on the console where it is possible to pass an input file (.yaml) 
to be consumed, this file defines the steps and validations of the end-to-end test.

![alt text](https://https://github.com/carloshjoaquim/E2Easy-Go/images/master/cmd.png?raw=true)



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