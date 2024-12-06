package org.akrck02.valhalla.core.sdk.model.user

import kotlinx.serialization.Serializable

/**
 * This class represents a user in the app
 */
@Serializable
data class User(
    var id: String? = null,
    var username: String? = null,
    var email: String? = null,
    var password: String? = null,
    var validated: Boolean? = null,
    var validationCode: String? = null,
    var profilePicturePath: String? = null,
    var creationTime: Long? = null
)
