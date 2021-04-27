import axios from '../../utils/axios'
import { ProfileDetailsJSONSchema, ProfileDetails } from './types'

type RegisterPayload = {
  name: string
  email: string
  password: string
}

type UpdateProfilePayload = {
  name: string
  bio: string
  location: string
  website: string
  birthDate: string
}

export async function registerAccount({
  name,
  email,
  password,
}: RegisterPayload) {
  try {
    await axios.post('/users/register', { name, email, password })
  } catch (error) {
    throw error
  }
}

function parseUpdateProfilePayload(payload: UpdateProfilePayload) {
  return {
    display_name: payload.name,
    bio: payload.bio,
    location: payload.location,
    website: payload.website,
    birth_date: payload.birthDate,
  }
}

export async function updateProfile(
  payload: UpdateProfilePayload
): Promise<void> {
  const jsonPayload = parseUpdateProfilePayload(payload)
  try {
    await axios.patch('/users/profile', jsonPayload)
  } catch (error) {
    throw error
  }
}

export async function followUser(userId: string) {
  try {
    await axios.post(`/relationships/follow/${userId}`)
  } catch (error) {
    throw error
  }
}

export async function unfollowUser(userId: string) {
  try {
    await axios.delete(`/relationships/unfollow/${userId}`)
  } catch (error) {
    throw error
  }
}

function parseProfileDetailsResponse(
  data: ProfileDetailsJSONSchema
): ProfileDetails {
  return {
    id: data.id,
    name: data.name,
    handle: data.handle,
    bio: data.bio,
    location: data.location,
    website: data.website,
    birthDate:
      data.birth_date !== '0001-01-01T00:00:00Z' ? data.birth_date : '',
    photoURL: data.photo_url,
    followersCount: data.followers_count,
    followingsCount: data.followings_count,
    isFollowing: data.is_following,
    joinedAt: data.joined_at,
  }
}

export async function fetchProfileDetails(
  handle: string
): Promise<ProfileDetails> {
  try {
    const { data } = await axios.get<ProfileDetailsJSONSchema>(
      `/users/${handle}`
    )

    return parseProfileDetailsResponse(data)
  } catch (error) {
    throw error
  }
}
