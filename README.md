# IoT2016 Bluemix + OpenWhisk Coding Tutorial

## Prereqs

We will do this tutorial using Golang, but every language that support MQTT and HTTP request will work.

### MQTT

The following link provides a list with all MQTT libraries:

https://github.com/mqtt/mqtt.github.io/wiki/libraries

**Most Common:**

- Google GO:

  https://github.com/eclipse/paho.mqtt.golang

- Java:

  https://eclipse.org/paho/clients/java/

- NodeJS (NPM):

  https://www.npmjs.com/package/mqtt

- JavaScript:

  https://www.npmjs.com/package/paho-mqtt

---

## OpenWhisk

We will trigger OpenWhisk using the REST API. So you can use your favorite http library for the programming language you want to use.

You can find the API documentation here

http://petstore.swagger.io/?url=https://raw.githubusercontent.com/openwhisk/openwhisk/master/core/controller/src/main/resources/whiskswagger.json

or here (i like this link more ;) )

https://console.ng.bluemix.net/apidocs/98?&language=node#get-all-namespaces-for-authenticated-user

The base URL for Bluemix is: `https://openwhisk.ng.bluemix.net/api/v1/`

Moreover, API calls need the following HTTP headers to set:

```
Content-Type: application/json
Authorization: Basic <token>
```

_Note: You will get the Base64 encoded token from the OpenWhisk UI. Alternatively, you can generate the token using the OpenWhisk auth-token._
