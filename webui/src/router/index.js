import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import UserProfileView from '../views/UserProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/:userId/profile', component: UserProfileView},
		//{path: '/link1', component: HomeView},
		//{path: '/link2', component: HomeView},
		//{path: '/some/:id/link', component: HomeView},
		//{path: '/'}
	]
})

export default router
