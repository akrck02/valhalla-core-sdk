package org.akrck02.valhalla.core.sdk.model.user

/**
 * This class represents the needed auth
 * parameters to access the API
 */
data class AuthLogin(
    val email: String,
    val authToken: String
)
