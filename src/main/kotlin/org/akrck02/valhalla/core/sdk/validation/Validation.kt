package org.akrck02.valhalla.core.sdk.validation

import org.akrck02.valhalla.core.sdk.error.ErrorCode
import org.akrck02.valhalla.core.sdk.model.exception.ServiceException
import org.akrck02.valhalla.core.sdk.model.http.HttpStatusCode

const val MINIMUM_CHARACTERS_FOR_PASSWORD = 16
const val MINIMUM_CHARACTERS_FOR_EMAIL = 5

/**
 * Validate string as an email to match the requirements
 * @throws ServiceException if the email is invalid
 */
@Suppress("unused")
fun String.validateEmail() {

    takeIf { it.isNotBlank() } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidRequest,
        message = "Email is empty."
    )

    takeIf { it.length >= MINIMUM_CHARACTERS_FOR_EMAIL } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidEmail,
        message = "Email must have at least $MINIMUM_CHARACTERS_FOR_EMAIL characters."
    )

    takeIf { it.contains("@") } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidEmail,
        message = "Email must have one @."
    )

    takeIf { it.contains(".") } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidEmail,
        message = "Email must have at least one dot."
    )

}

/**
 * Validate string as a password to match the requirements
 * @throws ServiceException if the password is invalid
 */
@Suppress("unused")
fun String.validatePassword() {

    takeIf { it.isNotBlank() } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidPassword,
        message = "Password is empty."
    )

    takeIf { it.length >= MINIMUM_CHARACTERS_FOR_PASSWORD } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidPassword,
        message = "Password must have at least $MINIMUM_CHARACTERS_FOR_PASSWORD characters."
    )

    takeIf { it.isLowercase().not().and(it.isUppercase().not()) } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidPassword,
        message = "Password must have at least one lowercase and uppercase character."
    )

    takeIf { it.containsSpecialCharacters() } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidPassword,
        message = "Password must have at least one special character."
    )

    takeIf { it.containsNumbers() } ?: throw ServiceException(
        status = HttpStatusCode.BadRequest,
        code = ErrorCode.InvalidPassword,
        message = "Password must have at least one number."
    )
}

/**
 * Get if a String is lowercase
 * @return true if it is, false otherwise
 */
fun String.isLowercase(): Boolean = this.lowercase() == this

/**
 * Get if a String is uppercase
 * @return true if it is, false otherwise
 */
fun String.isUppercase(): Boolean = this.uppercase() == this

/**
 * Get if a String contains numbers
 * @return true if it is, false otherwise
 */
fun String.containsNumbers(): Boolean = this.contains(Regex("[0-9]+"))

/**
 * Get if a String contains special characters
 * @return true if it is, false otherwise
 */
fun String.containsSpecialCharacters(): Boolean = this.contains(Regex("[¡!@#$%^&¿?*]+"))
