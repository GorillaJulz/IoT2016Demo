# IoT2016 Bluemix + OpenWhisk Coding Tutorial

## Prereqs

We will do this tutorial using Golang, but every language that support MQTT and HTTP requests will work.

How to install Golang:

https://golang.org/doc/install

### MQTT

The following link provides a list of all MQTT libraries:

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

**Bluemix MQTT Connections**

To setup an _Device_ connection you will need:

- **ClientID:** `d:$org:$type:$id`
- **MQTT Username:** `use-token-auth`
- **MQTT Password:** Auth-Token you get when setting up an new device
- **Connection URL:**

  ```
  tcp://<org>.messaging.internetofthings.ibmcloud.com:1883
  ```
- **Topic:**

  ```
  iot-2/evt/<event_id>/fmt/<format_string>
  ```

  - *event_id: [status|input|etc.]*
  - *format_string: [json|txt|xml|csv]*

To setup an _App_ connection you will need:

- **ClientID:** `a:$org:$AppID`
- **MQTT Username:** `Api-Key`
- **MQTT Password:** `Api Token`
- **Connection URL:** Same as for Devices
- **Topic:**

  ```
  iot-2/type/<device-type>/id/<device-id>/evt/<event_type>/fmt/<format_string>
  ```

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
