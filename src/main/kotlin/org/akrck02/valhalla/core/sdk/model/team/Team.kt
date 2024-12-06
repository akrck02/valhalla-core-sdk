package org.akrck02.valhalla.core.sdk.model.team

import kotlinx.serialization.Contextual
import kotlinx.serialization.SerialName
import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represents a team in the app
 */
@Suppress("unused")
data class Team(

    @SerialName("_id") @Contextual var id: String? = null,
    
    var name: String? = null,
    var owner: User? = null,
    var description: String? = null,
    var profilePicturePath: String? = null,

    var projects: MutableList<Int> = mutableListOf(),
    var members: MutableList<User> = mutableListOf(),

    @Contextual var creationTime: LocalDateTime? = null,
    @Contextual var updateTime: LocalDateTime? = null

)