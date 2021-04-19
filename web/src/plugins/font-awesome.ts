import { App } from '@vue/runtime-core'
import { library as fontAwesomeLibrary } from '@fortawesome/fontawesome-svg-core'
import {
  faPlus,
  faArrowLeft,
  faAngleDown,
  faCheck,
  faStar,
  faComment,
  faRetweet,
  faHeart,
  faShareSquare,
  faSearch,
  faCog,
  faHome,
  faHashtag,
  faBell,
  faEnvelope,
  faBookmark,
  faClipboardList,
  faUser,
  faEllipsisH,
  faTimes,
  faLink,
  faMapMarkerAlt,
  faBirthdayCake,
  faCalendarAlt,
} from '@fortawesome/free-solid-svg-icons'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

export default function () {
  fontAwesomeLibrary.add(
    faPlus,
    faArrowLeft,
    faAngleDown,
    faCheck,
    faStar,
    faComment,
    faRetweet,
    faHeart,
    faShareSquare,
    faSearch,
    faCog,
    faHome,
    faHashtag,
    faBell,
    faEnvelope,
    faBookmark,
    faClipboardList,
    faUser,
    faEllipsisH,
    faTwitter,
    faTimes,
    faLink,
    faMapMarkerAlt,
    faBirthdayCake,
    faCalendarAlt
  )

  return FontAwesomeIcon
}
