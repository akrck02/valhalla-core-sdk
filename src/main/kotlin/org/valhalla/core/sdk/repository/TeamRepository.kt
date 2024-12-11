package org.valhalla.core.sdk.repository

import org.valhalla.core.sdk.model.exception.ServiceException
import org.valhalla.core.sdk.model.team.Team

/** This interface represents the data access layer for the teams. */
@Suppress("unused")
interface TeamRepository {

    /**
     * Register a new [Team] returning the id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun register(team: Team?): String

    /**
     * Get a [Team] by id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun get(id: String?): Team

    /**
     * Get a [List] of the [Team]s a user belongs to by user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getAllByMember(userId: String?): List<Team>

    /**
     * Get a [List] of the [Team]s belonging to a user by user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getAllByOwner(userId: String?): List<Team>

    /**
     * Update the [Team] data given the team id and new team data.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun update(id: String?, team: Team?)

    /**
     * Update the owner of the [Team] by id and user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun updateOwner(id: String?, userId: String?)

    /**
     * Delete a [Team] by id
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun delete(id: String?)

    /**
     * Add a member to a team by team id and user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun addMember(id: String, userId: String?)

    /**
     * Delete a member from a team by team id and user id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun deleteMember(id: String, userId: String?)

    /**
     * Search all the [Team]s matching the name with the filter [String].
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun searchByName(filter: String?): List<Team>
}