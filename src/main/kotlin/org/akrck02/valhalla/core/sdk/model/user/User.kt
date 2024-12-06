package org.akrck02.valhalla.core.sdk.model.user

import kotlinx.serialization.Contextual
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import java.time.LocalDateTime

/**
 * This class represents a user in the app
 */
@Serializable
data class User(

    @SerialName("_id") @Contextual var id: String? = null,
    
    var username: String? = null,
    var email: String? = null,
    var password: String? = null,

    var validated: Boolean? = null,
    var validationCode: String? = null,

    var profilePicturePath: String? = null,

    @Contextual var creationTime: LocalDateTime? = null

)
