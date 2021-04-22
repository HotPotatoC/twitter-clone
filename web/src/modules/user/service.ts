import axios from '../../utils/axios'
import { ProfileDetailsJSONSchema, ProfileDetails } from './types'

interface RegisterPayload {
  name: string
  email: string
  password: string
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
    followersCount: data.followers_count,
    followingsCount: data.followings_count,
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
