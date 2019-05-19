<template>
	<div class="profile-wrapper">
        <!-- <div class="sidebar-wrapper"> 
            <ul> 
                <li v-for="item in menu" :key="item.id">
                    <a href=""> {{item.title}} </a>
                </li>
            </ul>
        </div> -->
        <div class="content-wrapper max-width">
            <div class="user-data">
                <div class="card">
                    <div class="card-body">
                        <h3 class="card-title"><span>Name: </span>{{userData.firstname}} {{userData.lastname}}</h3>
                        <h6 class="card-subtitle mb-2 text-muted"><span>Username: </span>{{userData.username}}</h6>
                        <p class="card-text"><span>E-mail: </span>{{userData.email}}</p>
                        <a href="#" class="card-link edit">Edit</a>
                    </div>
                </div>
            </div>
            <div class="educational">
                <div class="card">
                    <div class="card-body">
                        <h4 class="card-title">Oblasti u kojima vam treba dodatno obrazovanje i obuka: </h4>
                        <div class="chbx">
                            <label>{{userData.education_flags.e_edu_eu_standards.name}}</label> 
                            <input type="checkbox" :checked="userData.education_flags.e_edu_eu_standards.set">
                        </div>
                        <div class="chbx">
                            <label>{{userData.education_flags.e_edu_financial_reporting.name}}</label> 
                            <input type="checkbox" :checked="userData.education_flags.e_edu_financial_reporting.set">
                        </div>
                        
                        <div class="chbx">
                            <label>{{userData.education_flags.e_edu_market_emerge.name}}</label> 
                            <input type="checkbox" :checked="userData.education_flags.e_edu_market_emerge.set">
                        </div>

                        <div class="chbx">
                            <label>{{userData.education_flags.e_edu_project_mngmt.name}}</label> 
                            <input type="checkbox" :checked="userData.education_flags.e_edu_project_mngmt.set">
                        </div>

                        <div class="chbx">
                            <label>{{userData.education_flags.e_edu_writing_projects.name}}</label> 
                            <input type="checkbox" :checked="userData.education_flags.e_edu_writing_projects.set">
                        </div>                        

                        <div class="chbx">
                            <label>{{userData.experience_flags.e_enlish_fluent.name}}</label> 
                            <input type="checkbox" :checked="userData.experience_flags.e_enlish_fluent.set">
                        </div>
                        <div class="chbx">
                            <label>{{userData.experience_flags.e_eu_applied_before.name}}</label> 
                            <input type="checkbox" :checked="userData.experience_flags.e_eu_applied_before.set">
                        </div>
                        
                        <div class="chbx">
                            <label>{{userData.experience_flags.e_eu_funded_before.name}}</label> 
                            <input type="checkbox" :checked="userData.experience_flags.e_eu_funded_before.set">
                        </div>

                        <div class="chbx"   >
                            <label>{{userData.experience_flags.e_eu_subsidy_before.name}}</label> 
                            <input type="checkbox" :checked="userData.experience_flags.e_eu_subsidy_before.set">
                        </div>

                        <!-- <div>
                            <label>{{userData.experience_flags.e_wrote_projects.name}}</label> 
                            <input type="checkbox" :checked="userData.experience_flags.e_wrote_projects.set">
                        </div> -->
                    </div>
                </div>                
            </div>
            <div class="company">
                  <div class="card">
                    <div class="card-body" v-for="company in companies" :key="company.id">
                        <h3 class="card-title">Company: {{company.Name}} </h3>
                        <a href="#" class="card-link add">Add new +</a>
                    </div>
                </div>
            </div>
        </div>
        
	</div>
</template>

<script>

export default {
  name: 'Profile',
  components:{
  },
  data(){
      return{
           menu:[
              {"title":"Programi", "id":1},
              {"title":"Pozivi", "id":2},
              {"title":"Vesti", "id":3},
              {"title":"Dogadjaji", "id":4},
              {"title":"Pitanja", "id":5},
              {"title":"Znanje", "id":6},
              {"title":"Konsultanti", "id":7},
              {"title":"FAQ", "id":8}
          ],
          userData:{},
          userInfo:{},
          companies:{}
      }
  },
  components: {
    
  },
  mounted(){
    this.getUser()
    this.getCompanies()
  },
  methods:{
      getUser(){
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };

            var token = localStorage.getItem('token');
            var userId = this.$router.currentRoute.path.slice(6);
            this.axios(
                {method: "GET", "url": "http://localhost:8081/users/v1/user",
                "headers": { "Authorization": 'Bearer ' + token},
                config
                })
                .then(result => {
                this.userData = result.data;
                console.log(this.userData)
            }, error => {
                console.error(error);
            });;
      },
      getCompanies(){
            var config = {
                 headers: {'Access-Control-Allow-Origin': '*'}
            };
            var token = localStorage.getItem('token');
            var userId = this.$router.currentRoute.path.slice(6);
            this.axios(
                {method: "GET", "url": "http://localhost:8081/users/v1/user/" + userId + '/companies',
                "headers": { "Authorization": 'Bearer ' + token},
                config
                })
                .then(result => {
                this.companies = result.data.Companies;
                console.log(this.companies)
            }, error => {
                console.error(error);
            });;
      }
  }
}

</script>

<style lang="scss" scoped>
.max-width{
    width:1240px;
    margin:0 auto;
}
.profile-wrapper{
    min-height:100vh;
    background-color:#f4f4f4;
    .content-wrapper{
            .card{
                justify-self: flex-start;
                align-items:flex-start;
                border:none;
                text-align:left;
                border-top:1px solid #e6e6e6;
                .card-body{
                    width:100%;
                    padding:40px 20px;
                }
                h3,p,h6{
                    font-weight: 900;
                }
                span{
                    font-weight:500;
                }
            }
            .educational{
                .chbx{
                    display: flex;
                    justify-content: space-between;
                }
            }
            .company{
                .card-body{
                    display: flex;
                    justify-content: space-between;
                    align-items: flex-start;
                    .add{
                    background-color: #3483ca;
                    padding:8px 24px;
                    color:white;
                    border-radius: 8px;
                    font-weight: 900;
                    float: right;
                    }
                    width:100%;
                }
            }
        .user-data{
            padding: 40px 0px;
            background: white;
            .card{
                justify-self: flex-start;
                -webkit-box-align: start;
                -ms-flex-align: start;
                align-items: flex-start;
                text-align: left;
                background: none;
                border: none;
                .card-title{
                    margin-bottom:2rem;
                }
                .card-body{
                    width:100%;
                }
                .edit{
                    float: right;
                    background-color: #3483ca;
                    padding:8px 24px;
                    color:white;
                    border-radius: 8px;
                }
            }
        }
    }
}

</style>
