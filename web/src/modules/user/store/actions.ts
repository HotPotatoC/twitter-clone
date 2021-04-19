import { ActionTree } from 'vuex'
import { AugmentedActionContext } from '../../../store/types'
import { Mutations, MutationTypes } from './mutations'
import { ProfileDetails, State } from './state'

import axios from '../../../services/axios'

export enum ActionTypes {
  GET_PROFILE_DETAILS = 'GET_PROFILE_DETAILS',
}

export interface Actions {
  [ActionTypes.GET_PROFILE_DETAILS](
    { commit }: AugmentedActionContext<Mutations, State>,
    username: string | string[]
  ): Promise<void>
}

interface GetProfileDetailsJSON {
  id: number
  name: string
  bio: string
  location: string
  website: string
  birth_date: string
  followers_count: number
  followings_count: number
  joined_at: string
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.GET_PROFILE_DETAILS]({ commit }, username): Promise<void> {
    try {
      const response = await axios.get<GetProfileDetailsJSON>(
        `/users/${username}`
      )

      const profileDetails: ProfileDetails = {
        id: response.data.id,
        name: response.data.name,
        bio: response.data.bio,
        location: response.data.location,
        website: response.data.website,
        birthDate:
          response.data.birth_date !== '0001-01-01T00:00:00Z'
            ? response.data.birth_date
            : '',
        followersCount: response.data.followers_count,
        followingsCount: response.data.followings_count,
        joinedAt: response.data.joined_at,
      }

      commit(MutationTypes.SET_PROFILE_DETAILS, profileDetails)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: response.status,
        message: '',
      })
    } catch (error) {
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
}
