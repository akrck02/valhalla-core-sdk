package org.akrck02.valhalla.core.sdk.repository

import org.akrck02.valhalla.core.sdk.model.team.Team

/**
 * This interface represents the data access layer
 * for the team collection.
 */
@Suppress("unused")
interface TeamRepository {

    /**
     * Register a new team
     * @param team The team to register
     */
    fun register(team: Team?)

    /**
     * Get a team by id
     * @param id The team id
     * @return The requested team
     */
    fun get(id : String?) : Team ?

    /**
     * Get all the teams a user belongs to
     * @param userId The user id
     * @return the list of teams a user belongs to
     */
    fun getAllByMember(userId : String?) : List<Team>

    /**
     * Get all the teams belonging to a user
     * @param userId The user id
     * @return the list of teams belonging to a user
     */
    fun getAllByOwner(userId : String?) : List<Team>

    /**
     * Update the team data
     * @param id The team id
     * @param team The new team data
     */
    fun update(id : String?, team: Team?)

    /**
     * Update the owner of the team
     * @param id The team id
     * @param userId The new owner user id
     */
    fun updateOwner(id: String?, userId : String?)

    /**
     * Delete a team
     * @param id The team id
     */
    fun delete(id : String?)

    /**
     * Add a member to a team
     * @param id The team id
     * @param userId The new member user id
     */
    fun addMember(id : String , userId : String?)

    /**
     * Delete a member from a team
     * @param id The team id
     * @param userId The member user id
     */
    fun deleteMember(id : String , userId : String?)

    /**
     * Search all the teams matching the name with the filter string
     */
    fun searchByName(filter : String ?) : List<Team>
}