package org.akrck02.valhalla.core.sdk.model.exception

import org.akrck02.valhalla.core.sdk.error.ErrorCode
import org.akrck02.valhalla.core.sdk.model.http.HttpStatusCode

/**
 * This class represents an exception in the service
 */
@Suppress("unused")
class ServiceException(
    val status: HttpStatusCode = HttpStatusCode.InternalServerError,
    val code: ErrorCode? = ErrorCode.UnexpectedError,
    override val message: String = "Unexpected error"
) : Exception(message)
