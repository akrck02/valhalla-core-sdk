package org.akrck02.valhalla.core.sdk.model.project

import kotlinx.serialization.Contextual
import kotlinx.serialization.SerialName
import org.akrck02.valhalla.core.sdk.model.team.Team
import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represent a project in the app
 */
@Suppress("unused")
data class Project(

    @SerialName("_id") @Contextual var id: String? = null,
    
    var name: String? = null,
    var description: String? = null,
    var owner: User? = null,

    var teams: MutableList<Team> = mutableListOf(),

    @Contextual var creationTime: LocalDateTime? = null,
    @Contextual var updateTime: LocalDateTime? = null

)