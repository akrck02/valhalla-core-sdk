package org.akrck02.valhalla.core.sdk.model.team

import kotlinx.serialization.Serializable
import org.akrck02.valhalla.core.sdk.model.user.User
import org.bson.BsonType
import org.bson.codecs.pojo.annotations.BsonId
import org.bson.codecs.pojo.annotations.BsonRepresentation

/**
 * This class represents a team in the app
 */
@Suppress("unused")
@Serializable
data class Team(

    @BsonId
    @BsonRepresentation(BsonType.OBJECT_ID)
    var id: String? = null,
    var name: String? = null,
    var owner: User? = null,
    var description: String? = null,
    var profilePicturePath: String? = null,
    var projects: MutableList<Int> = mutableListOf(),
    var members: MutableList<User> = mutableListOf(),
    var creationTime: Long? = null,
    var updateTime: Long? = null
)