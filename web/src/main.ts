import { createApp } from 'vue'
import App from './App.vue'
import { router } from './routes'
import { store } from './store'

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
} from '@fortawesome/free-solid-svg-icons'
import { faTwitter } from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import './assets/tailwind.css'

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
  faTwitter
)

const app = createApp(App)

app.component('font-awesome', FontAwesomeIcon)

app.use(store)
app.use(router)
app.mount('#app')
