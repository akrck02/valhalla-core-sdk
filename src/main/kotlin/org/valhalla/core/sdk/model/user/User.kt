package org.valhalla.core.sdk.model.user

import kotlinx.serialization.Serializable
import org.bson.BsonType
import org.bson.codecs.pojo.annotations.BsonId
import org.bson.codecs.pojo.annotations.BsonRepresentation
import org.valhalla.core.sdk.model.device.Device

/** This class represents a user in the app. */
@Serializable
data class User(
    @BsonId @BsonRepresentation(BsonType.OBJECT_ID) var id: String? = null,
    var username: String? = null,
    var email: String? = null,
    var password: String? = null,
    var validated: Boolean? = null,
    var validationCode: String? = null,
    var profilePicturePath: String? = null,
    var creationTime: Long? = null,
    var devices: MutableList<Device> = mutableListOf()
)
