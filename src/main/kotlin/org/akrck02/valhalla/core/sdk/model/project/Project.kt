package org.akrck02.valhalla.core.sdk.model.project

import kotlinx.serialization.Serializable
import org.akrck02.valhalla.core.sdk.model.team.Team
import org.akrck02.valhalla.core.sdk.model.user.User
import org.bson.BsonType
import org.bson.codecs.pojo.annotations.BsonId
import org.bson.codecs.pojo.annotations.BsonRepresentation

/**
 * This class represent a project in the app
 */
@Suppress("unused")
@Serializable
data class Project(

    @BsonId
    @BsonRepresentation(BsonType.OBJECT_ID)
    var id: String? = null,
    var name: String? = null,
    var description: String? = null,
    var owner: User? = null,
    var teams: MutableList<Team> = mutableListOf(),
    var creationTime: Long? = null,
    var updateTime: Long? = null
)

