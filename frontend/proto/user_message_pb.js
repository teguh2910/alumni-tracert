// source: user_message.proto
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
var alumni_appraiser_message_pb = require('./alumni_appraiser_message_pb.js');
goog.object.extend(proto, alumni_appraiser_message_pb);
var user_answer_message_pb = require('./user_answer_message_pb.js');
goog.object.extend(proto, user_answer_message_pb);
var generic_message_pb = require('./generic_message_pb.js');
goog.object.extend(proto, generic_message_pb);
goog.exportSymbol('proto.proto.ChangePassword', null, global);
goog.exportSymbol('proto.proto.LoginInput', null, global);
goog.exportSymbol('proto.proto.User', null, global);
goog.exportSymbol('proto.proto.UserListResponse', null, global);
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
proto.proto.User = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.proto.User.repeatedFields_, null);
};
goog.inherits(proto.proto.User, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.User.displayName = 'proto.proto.User';
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
proto.proto.UserListResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.UserListResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.UserListResponse.displayName = 'proto.proto.UserListResponse';
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
proto.proto.ChangePassword = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.ChangePassword, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.ChangePassword.displayName = 'proto.proto.ChangePassword';
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
proto.proto.LoginInput = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.proto.LoginInput, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.proto.LoginInput.displayName = 'proto.proto.LoginInput';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.proto.User.repeatedFields_ = [8];



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
proto.proto.User.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.User.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.User} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.User.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    email: jspb.Message.getFieldWithDefault(msg, 2, ""),
    name: jspb.Message.getFieldWithDefault(msg, 3, ""),
    isActived: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    userType: jspb.Message.getFieldWithDefault(msg, 5, 0),
    alumni: (f = msg.getAlumni()) && alumni_message_pb.Alumni.toObject(includeInstance, f),
    alumniAppraiser: (f = msg.getAlumniAppraiser()) && alumni_appraiser_message_pb.AlumniAppraiser.toObject(includeInstance, f),
    userAnswerList: jspb.Message.toObjectList(msg.getUserAnswerList(),
    user_answer_message_pb.UserAnswer.toObject, includeInstance),
    created: jspb.Message.getFieldWithDefault(msg, 9, ""),
    updated: jspb.Message.getFieldWithDefault(msg, 10, ""),
    token: jspb.Message.getFieldWithDefault(msg, 11, "")
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
 * @return {!proto.proto.User}
 */
proto.proto.User.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.User;
  return proto.proto.User.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.User} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.User}
 */
proto.proto.User.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsActived(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setUserType(value);
      break;
    case 6:
      var value = new alumni_message_pb.Alumni;
      reader.readMessage(value,alumni_message_pb.Alumni.deserializeBinaryFromReader);
      msg.setAlumni(value);
      break;
    case 7:
      var value = new alumni_appraiser_message_pb.AlumniAppraiser;
      reader.readMessage(value,alumni_appraiser_message_pb.AlumniAppraiser.deserializeBinaryFromReader);
      msg.setAlumniAppraiser(value);
      break;
    case 8:
      var value = new user_answer_message_pb.UserAnswer;
      reader.readMessage(value,user_answer_message_pb.UserAnswer.deserializeBinaryFromReader);
      msg.addUserAnswer(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreated(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdated(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
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
proto.proto.User.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.User.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.User} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.User.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getIsActived();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getUserType();
  if (f !== 0) {
    writer.writeUint32(
      5,
      f
    );
  }
  f = message.getAlumni();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      alumni_message_pb.Alumni.serializeBinaryToWriter
    );
  }
  f = message.getAlumniAppraiser();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      alumni_appraiser_message_pb.AlumniAppraiser.serializeBinaryToWriter
    );
  }
  f = message.getUserAnswerList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      user_answer_message_pb.UserAnswer.serializeBinaryToWriter
    );
  }
  f = message.getCreated();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getUpdated();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
};


/**
 * optional uint64 id = 1;
 * @return {number}
 */
proto.proto.User.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string email = 2;
 * @return {string}
 */
proto.proto.User.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setEmail = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.proto.User.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional bool is_actived = 4;
 * @return {boolean}
 */
proto.proto.User.prototype.getIsActived = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setIsActived = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional uint32 user_type = 5;
 * @return {number}
 */
proto.proto.User.prototype.getUserType = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setUserType = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional Alumni alumni = 6;
 * @return {?proto.proto.Alumni}
 */
proto.proto.User.prototype.getAlumni = function() {
  return /** @type{?proto.proto.Alumni} */ (
    jspb.Message.getWrapperField(this, alumni_message_pb.Alumni, 6));
};


/**
 * @param {?proto.proto.Alumni|undefined} value
 * @return {!proto.proto.User} returns this
*/
proto.proto.User.prototype.setAlumni = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.clearAlumni = function() {
  return this.setAlumni(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.User.prototype.hasAlumni = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional AlumniAppraiser alumni_appraiser = 7;
 * @return {?proto.proto.AlumniAppraiser}
 */
proto.proto.User.prototype.getAlumniAppraiser = function() {
  return /** @type{?proto.proto.AlumniAppraiser} */ (
    jspb.Message.getWrapperField(this, alumni_appraiser_message_pb.AlumniAppraiser, 7));
};


/**
 * @param {?proto.proto.AlumniAppraiser|undefined} value
 * @return {!proto.proto.User} returns this
*/
proto.proto.User.prototype.setAlumniAppraiser = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.clearAlumniAppraiser = function() {
  return this.setAlumniAppraiser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.User.prototype.hasAlumniAppraiser = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * repeated UserAnswer user_answer = 8;
 * @return {!Array<!proto.proto.UserAnswer>}
 */
proto.proto.User.prototype.getUserAnswerList = function() {
  return /** @type{!Array<!proto.proto.UserAnswer>} */ (
    jspb.Message.getRepeatedWrapperField(this, user_answer_message_pb.UserAnswer, 8));
};


/**
 * @param {!Array<!proto.proto.UserAnswer>} value
 * @return {!proto.proto.User} returns this
*/
proto.proto.User.prototype.setUserAnswerList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.proto.UserAnswer=} opt_value
 * @param {number=} opt_index
 * @return {!proto.proto.UserAnswer}
 */
proto.proto.User.prototype.addUserAnswer = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.proto.UserAnswer, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.clearUserAnswerList = function() {
  return this.setUserAnswerList([]);
};


/**
 * optional string created = 9;
 * @return {string}
 */
proto.proto.User.prototype.getCreated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setCreated = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string updated = 10;
 * @return {string}
 */
proto.proto.User.prototype.getUpdated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setUpdated = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string token = 11;
 * @return {string}
 */
proto.proto.User.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.User} returns this
 */
proto.proto.User.prototype.setToken = function(value) {
  return jspb.Message.setProto3StringField(this, 11, value);
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
proto.proto.UserListResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.UserListResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.UserListResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.UserListResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    user: (f = msg.getUser()) && proto.proto.User.toObject(includeInstance, f),
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
 * @return {!proto.proto.UserListResponse}
 */
proto.proto.UserListResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.UserListResponse;
  return proto.proto.UserListResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.UserListResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.UserListResponse}
 */
proto.proto.UserListResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.proto.User;
      reader.readMessage(value,proto.proto.User.deserializeBinaryFromReader);
      msg.setUser(value);
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
proto.proto.UserListResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.UserListResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.UserListResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.UserListResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.proto.User.serializeBinaryToWriter
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
 * optional User user = 1;
 * @return {?proto.proto.User}
 */
proto.proto.UserListResponse.prototype.getUser = function() {
  return /** @type{?proto.proto.User} */ (
    jspb.Message.getWrapperField(this, proto.proto.User, 1));
};


/**
 * @param {?proto.proto.User|undefined} value
 * @return {!proto.proto.UserListResponse} returns this
*/
proto.proto.UserListResponse.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.UserListResponse} returns this
 */
proto.proto.UserListResponse.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.UserListResponse.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional ListInput list_input = 2;
 * @return {?proto.proto.ListInput}
 */
proto.proto.UserListResponse.prototype.getListInput = function() {
  return /** @type{?proto.proto.ListInput} */ (
    jspb.Message.getWrapperField(this, generic_message_pb.ListInput, 2));
};


/**
 * @param {?proto.proto.ListInput|undefined} value
 * @return {!proto.proto.UserListResponse} returns this
*/
proto.proto.UserListResponse.prototype.setListInput = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.proto.UserListResponse} returns this
 */
proto.proto.UserListResponse.prototype.clearListInput = function() {
  return this.setListInput(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.proto.UserListResponse.prototype.hasListInput = function() {
  return jspb.Message.getField(this, 2) != null;
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
proto.proto.ChangePassword.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.ChangePassword.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.ChangePassword} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.ChangePassword.toObject = function(includeInstance, msg) {
  var f, obj = {
    newPassword: jspb.Message.getFieldWithDefault(msg, 1, ""),
    reNewPassword: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.proto.ChangePassword}
 */
proto.proto.ChangePassword.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.ChangePassword;
  return proto.proto.ChangePassword.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.ChangePassword} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.ChangePassword}
 */
proto.proto.ChangePassword.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewPassword(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setReNewPassword(value);
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
proto.proto.ChangePassword.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.ChangePassword.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.ChangePassword} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.ChangePassword.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNewPassword();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getReNewPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string new_password = 1;
 * @return {string}
 */
proto.proto.ChangePassword.prototype.getNewPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.ChangePassword} returns this
 */
proto.proto.ChangePassword.prototype.setNewPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string re_new_password = 2;
 * @return {string}
 */
proto.proto.ChangePassword.prototype.getReNewPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.ChangePassword} returns this
 */
proto.proto.ChangePassword.prototype.setReNewPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
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
proto.proto.LoginInput.prototype.toObject = function(opt_includeInstance) {
  return proto.proto.LoginInput.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.proto.LoginInput} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.LoginInput.toObject = function(includeInstance, msg) {
  var f, obj = {
    email: jspb.Message.getFieldWithDefault(msg, 1, ""),
    password: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.proto.LoginInput}
 */
proto.proto.LoginInput.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.proto.LoginInput;
  return proto.proto.LoginInput.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.proto.LoginInput} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.proto.LoginInput}
 */
proto.proto.LoginInput.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
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
proto.proto.LoginInput.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.proto.LoginInput.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.proto.LoginInput} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.proto.LoginInput.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string email = 1;
 * @return {string}
 */
proto.proto.LoginInput.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.LoginInput} returns this
 */
proto.proto.LoginInput.prototype.setEmail = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.proto.LoginInput.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.proto.LoginInput} returns this
 */
proto.proto.LoginInput.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.proto);
