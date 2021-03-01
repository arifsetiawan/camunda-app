<template>
  <div>
    <CRow>
      <CCol lg="12">
        <CCard v-if="variables">
            <CCardHeader>
                <strong> Leave Request Approval - {{variables.leaveId.value}}</strong>
            </CCardHeader>
          <CCardBody>
            <CInput
              label="User ID"
              placeholder="User ID"
              readonly
              v-model="variables.id.value" 
            />
            <CInput
              label="Name"
              placeholder="Name"
              readonly
              v-model="variables.name.value"
            />
            <CInput
              label="Email"
              placeholder="Email"
              readonly
              v-model="variables.email.value" 
            />
            <CInput
              label="Employee Role"
              placeholder="Role"
              readonly
              v-model="variables.level.value" 
            />
            <CInput
              label="Leave Days"
              placeholder="Number of leave days"
              readonly
              v-model="variables.days.value" 
            />
            <CSelect
                label="Action"
                :options=options
                placeholder="Please select"
                :value.sync="selectedAction"
            />
            <br>
            <CRow>
                <CCol col="12" class="text-right">
                    <CButton color="primary" class="px-4" @click="submit">Submit</CButton>
                </CCol>
            </CRow>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LeaveRequestApproval',
  data() {
    return {
      variables: null,
      options: ['Approve', 'Reject'],
      selectedAction: ""
    };
  },
  computed: {
  },
  methods: {
    submit: function() {
      if (this.selectedAction != "") {
        var data = {
          variables: {
            approved: {
              value: this.selectedAction === 'Approve',
              type: "Boolean"
            }
          }
        }

        axios.post(process.env.VUE_APP_CAMUNDA_PROXY_URL + '/tasks/' + this.$route.params.taskId + '/complete', data)
        .then(resp => {
          console.log(resp)
        })
        .catch(err => {
          console.log(err);
        })      
      }
    }
  },
  created() {
    axios
      .get(process.env.VUE_APP_CAMUNDA_PROXY_URL + '/tasks/' + this.$route.params.taskId)
      .then(resp => {
        console.log(resp.data)
        this.variables = resp.data
      })
  }
}
</script>
