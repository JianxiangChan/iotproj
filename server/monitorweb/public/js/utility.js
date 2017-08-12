// Create a client instance
client = null;
connected = false;
var hostname,port,path,user,pass,keepAlive,timeout,ssl,cleanSession,lastWillTopic,lastWillQos,lastWillRetain,lastWillMessage,topic;

function setMqttEntity(mqttEntity){
	 console.log(mqttEntity);
    topic = mqttEntity.MqttTopic;
	hostname = mqttEntity.MqttHostname;
	port = mqttEntity.MqttPort;
	path = mqttEntity.MqttPath;
	user = mqttEntity.MqttUser;
	pass = mqttEntity.MqttPass;
	keepAlive = mqttEntity.MqttKeepAlive;
	timeout = mqttEntity.MqttTimeout;
	ssl = mqttEntity.MqttSsl;
	cleanSession = mqttEntity.MqttCleanSession;
	lastWillTopic = mqttEntity.MqttLastWillTopic;
	lastWillQos = mqttEntity.MqttLastWillQos;
	lastWillRetain = mqttEntity.MqttLastWillRetain;
	lastWillMessage = mqttEntity.MqttLastWillMessage;
	 console.log("Set mqtt entity success");
}

function connectionToggle() {
    var clientId =  ""+ new Date().getTime();
    if (connected) {
        disconnect(clientId);
    } else {
        connect(clientId);
    }
}

// called when the client connects
function onConnect(context) {
    console.log("Connected success");
    brfore = new Date().getTime();
    subscribe();
    connected = true;
}

function onFail(context) {
    console.log("Failed to connect" + context.errorMessage);
    connected = false;
    connectionToggle();
}

// called when the client loses its connection
function onConnectionLost(responseObject) {
    console.log("connect lost! " + responseObject);

    $.each(responseObject, function(p, v) {
        console.log(p + "-------" + v);

    });
    connected = false;
    connectionToggle();
}

// called when a message arrives
function onMessageArrived(message) {
    //	alert('Message Recieved: Topic: '+ message.destinationNam+ '. Payload: '+ message.payloadString+ '. QoS: '+ message.qos);
    var rtdata = $.parseJSON(message.payloadString);
    console.log(rtdata);
    setTimeout(function() {
        console.log(rtdata);
        rtupdate(rtdata)
    }, 500);
}

function connect(clientId) {

    if (path.length > 0) {
        client = new Paho.MQTT.Client(hostname, Number(port), path, clientId);
    } else {
        client = new Paho.MQTT.Client(hostname, Number(port), clientId);
    }
    //    alert('Connecting to Server: Hostname: '+ hostname+ '. Port: '+ port+ '. Path: '+ client.path+ '. Client ID: '+ clientId);

    // set callback handlers
    client.onConnectionLost = onConnectionLost;
    client.onMessageArrived = onMessageArrived;

    var options = {
        invocationContext: { host: hostname, port: port, path: client.path, clientId: clientId },
        timeout: timeout,
        keepAliveInterval: keepAlive,
        cleanSession: cleanSession,
        useSSL: ssl,
        onSuccess: onConnect,
        onFailure: onFail
    };

    if (user.length > 0) {
        options.userName = user;
    }

    if (pass.length > 0) {
        options.password = pass;
    }

    if (lastWillTopic.length > 0) {
        var lastWillMessage = new Paho.MQTT.Message(lastWillMessage);
        lastWillMessage.destinationName = lastWillTopic;
        lastWillMessage.qos = lastWillQos;
        lastWillMessage.retained = lastWillRetain;
        options.willMessage = lastWillMessage;
    }
    client.connect(options);
}

function disconnect() {
    alert('Disconnecting from Server');
    client.disconnect();
    connected = false;
    connectionToggle();
}

function subscribe() {
    var qos = 2;
    client.subscribe(topic, { qos: Number(qos) });
    //    alert('Subscribing to: Topic: '+ topic+ '. QoS: '+ qos);
}

function unsubscribe(subDevEUI) {
    //    alert('Unsubscribing from '+ topic);
    client.unsubscribe(topic, {
        onSuccess: unsubscribeSuccess,
        onFailure: unsubscribeFailure,
        invocationContext: { topic: topic }
    });
}

function unsubscribeSuccess(context) {
    //    alert('Successfully unsubscribed from '+ context.invocationContext.topic);
}

function unsubscribeFailure(context) {
    //	alert('Failed to  unsubscribe from '+ context.invocationContext.topic);
}

function safe_tags_regex(str) {
    return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
}