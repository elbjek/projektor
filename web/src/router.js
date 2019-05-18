  
import VueRouter from 'vue-router';

import login from './components/Login.vue'
// import userpage from './components/userpage.vue'
import landing from './components/LandingPage.vue'

export default new VueRouter({
    mode: 'history',
    routes: [

        {path:'/',component:landing},
        {path:'/login', component:login},
        // {path:'/userpage', component:userpage}
    ]
});