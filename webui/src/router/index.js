import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import UserProfileView from '../views/UserProfileView.vue'
import PostedImageView from '../views/PostedImageView.vue'
import HomeScreenView from '../views/HomeScreenView.vue'
import VisitingUserView from '../views/VisitingUserView.vue'
import VisitingPostView from '../views/VisitingPostView.vue'
import BlockedUserView from '../views/BlockedUserView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/:userId/profile', component: UserProfileView},
		{path: '/:userId/profile/:photoId', component: PostedImageView},
		{path: '/:userId/home', component: HomeScreenView},
		{path: '/:userId/profile/iteract/:userId', component: VisitingUserView},
		{path: '/:userId/profile/iteract/:userId/:photoId', component: VisitingPostView},
		{path: '/:userId/blocked', component: BlockedUserView},
		//{path: '/link1', component: HomeView},
		//{path: '/link2', component: HomeView},
		//{path: '/some/:id/link', component: HomeView},
		//{path: '/'}
	]
})

export default router
