  
import VueRouter from 'vue-router';

import login from './components/Login.vue'
import user from './components/Profile.vue'
import landing from './components/LandingPage.vue'
import home from './components/Home.vue'
import calls from './components/Calls.vue'
import call from './components/singleCall.vue'

export default new VueRouter({
    mode: 'history',
    routes: [

        {path:'/',component:landing},
        {path:'/login', component:login},
        {path:'/user/:id', component:user, props:true},
        {path:'/home', component:home},
        {path:'/calls', component:calls},
        {path:'/calls/:id', component:call}
    ]
});