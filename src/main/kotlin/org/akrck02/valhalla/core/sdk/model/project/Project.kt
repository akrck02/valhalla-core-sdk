package org.akrck02.valhalla.core.sdk.model.project

import kotlinx.serialization.Contextual
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import org.akrck02.valhalla.core.sdk.model.team.Team
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This class represent a project in the app
 */
@Suppress("unused")
@Serializable
data class Project(
    @Contextual
    @SerialName("_id") var id: String? = null,
    var name: String? = null,
    var description: String? = null,
    var owner: User? = null,
    var teams: MutableList<Team> = mutableListOf(),
    var creationTime: Long? = null,
    var updateTime: Long? = null
)