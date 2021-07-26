// source: alumni_appraiser_message.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var alumni_message_pb = require('./alumni_message_pb.js');
goog.object.extend(proto, alumni_message_pb);
var generic_message_pb = require('./generic_message_pb.js');
goog.object.extend(proto, generic_message_pb);
goog.exportSymbol('proto.proto.AlumniAppraiser', null, global);
goog.exportSymbol('proto.proto.AlumniAppraiserListResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.AlumniAppraiser = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.AlumniAppraiser, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.AlumniAppraiser.displayName = 'proto.proto.AlumniAppraiser';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.proto.AlumniAppraiserListResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.AlumniAppraiserListResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.AlumniAppraiserListResponse.displayName = 'proto.proto.AlumniAppraiserListResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.AlumniAppraiser.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.AlumniAppraiser.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.AlumniAppraiser} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.AlumniAppraiser.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    userId: jspb.Message.getFieldWithDefault(msg, 2, 0),
    alumni: (f = msg.getAlumni()) && alumni_message_pb.Alumni.toObject(includeInstance, f),
    name: jspb.Message.getFieldWithDefault(msg, 4, ""),
    instansi: jspb.Message.getFieldWithDefault(msg, 5, ""),
    position: jspb.Message.getFieldWithDefault(msg, 6, ""),
    alumniPosition: jspb.Message.getFieldWithDefault(msg, 7, ""),
    created: jspb.Message.getFieldWithDefault(msg, 8, ""),
    updated: jspb.Message.getFieldWithDefault(msg, 9, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.AlumniAppraiser}
 */
proto.proto.AlumniAppraiser.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.AlumniAppraiser;
  return proto.proto.AlumniAppraiser.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.AlumniAppraiser} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.AlumniAppraiser}
 */
proto.proto.AlumniAppraiser.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setUserId(value);
      break;
    case 3:
      var value = new alumni_message_pb.Alumni;
      reader.readMessage(value,alumni_message_pb.Alumni.deserializeBinaryFromReader);
      msg.setAlumni(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setInstansi(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPosition(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setAlumniPosition(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreated(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdated(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.AlumniAppraiser.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.AlumniAppraiser.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.AlumniAppraiser} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.AlumniAppraiser.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f !== 0) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getAlumni();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      alumni_message_pb.Alumni.serializeBinaryToWriter
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getInstansi();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPosition();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getAlumniPosition();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getCreated();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getUpdated();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.proto.AlumniAppraiser.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional uint64 user_id = 2;
 * @return {number}
 */
proto.proto.AlumniAppraiser.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional Alumni alumni = 3;
 * @return {?proto.proto.Alumni}
 */
proto.proto.AlumniAppraiser.prototype.getAlumni = function() {
  return /** @type{?proto.proto.Alumni} */ (
    jspb.Message.getWrapperField(this, alumni_message_pb.Alumni, 3));
};


/**
 * @param {?proto.proto.Alumni|undefined} value
 * @return {!proto.proto.AlumniAppraiser} returns this
*/
proto.proto.AlumniAppraiser.prototype.setAlumni = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.clearAlumni = function() {
  return this.setAlumni(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.AlumniAppraiser.prototype.hasAlumni = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string name = 4;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string instansi = 5;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getInstansi = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setInstansi = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string position = 6;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getPosition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setPosition = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string alumni_position = 7;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getAlumniPosition = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setAlumniPosition = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string created = 8;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getCreated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setCreated = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string updated = 9;
 * @return {string}
 */
proto.proto.AlumniAppraiser.prototype.getUpdated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.AlumniAppraiser} returns this
 */
proto.proto.AlumniAppraiser.prototype.setUpdated = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.proto.AlumniAppraiserListResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.AlumniAppraiserListResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.AlumniAppraiserListResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.AlumniAppraiserListResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    alumniAppraiser: (f = msg.getAlumniAppraiser()) && proto.proto.AlumniAppraiser.toObject(includeInstance, f),
    listInput: (f = msg.getListInput()) && generic_message_pb.ListInput.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.proto.AlumniAppraiserListResponse}
 */
proto.proto.AlumniAppraiserListResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.AlumniAppraiserListResponse;
  return proto.proto.AlumniAppraiserListResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.AlumniAppraiserListResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.AlumniAppraiserListResponse}
 */
proto.proto.AlumniAppraiserListResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.proto.AlumniAppraiser;
      reader.readMessage(value,proto.proto.AlumniAppraiser.deserializeBinaryFromReader);
      msg.setAlumniAppraiser(value);
      break;
    case 2:
      var value = new generic_message_pb.ListInput;
      reader.readMessage(value,generic_message_pb.ListInput.deserializeBinaryFromReader);
      msg.setListInput(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.proto.AlumniAppraiserListResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.AlumniAppraiserListResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.AlumniAppraiserListResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.AlumniAppraiserListResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAlumniAppraiser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.proto.AlumniAppraiser.serializeBinaryToWriter
    );
  }
  f = message.getListInput();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      generic_message_pb.ListInput.serializeBinaryToWriter
    );
  }
};


/**
 * optional AlumniAppraiser alumni_appraiser = 1;
 * @return {?proto.proto.AlumniAppraiser}
 */
proto.proto.AlumniAppraiserListResponse.prototype.getAlumniAppraiser = function() {
  return /** @type{?proto.proto.AlumniAppraiser} */ (
    jspb.Message.getWrapperField(this, proto.proto.AlumniAppraiser, 1));
};


/**
 * @param {?proto.proto.AlumniAppraiser|undefined} value
 * @return {!proto.proto.AlumniAppraiserListResponse} returns this
*/
proto.proto.AlumniAppraiserListResponse.prototype.setAlumniAppraiser = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.AlumniAppraiserListResponse} returns this
 */
proto.proto.AlumniAppraiserListResponse.prototype.clearAlumniAppraiser = function() {
  return this.setAlumniAppraiser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.AlumniAppraiserListResponse.prototype.hasAlumniAppraiser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional ListInput list_input = 2;
 * @return {?proto.proto.ListInput}
 */
proto.proto.AlumniAppraiserListResponse.prototype.getListInput = function() {
  return /** @type{?proto.proto.ListInput} */ (
    jspb.Message.getWrapperField(this, generic_message_pb.ListInput, 2));
};


/**
 * @param {?proto.proto.ListInput|undefined} value
 * @return {!proto.proto.AlumniAppraiserListResponse} returns this
*/
proto.proto.AlumniAppraiserListResponse.prototype.setListInput = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.AlumniAppraiserListResponse} returns this
 */
proto.proto.AlumniAppraiserListResponse.prototype.clearListInput = function() {
  return this.setListInput(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.AlumniAppraiserListResponse.prototype.hasListInput = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.proto);
