<template>
  <div>
    <CRow>
      <CCol sm="12">
        <CCard>
          <CCardBody>
            <CDataTable
              :items="tasks"
              tableFilter
              pagination
              hover
              striped
              sorter
            >
            <template #id="{item}">
              <td>
                <CLink :to="taskDetailLink(item.id)">{{item.id}}</CLink>
              </td>
            </template>
            </CDataTable>
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Tasks',
  data() {
    return {
      tasks: []
    };
  },
  methods: {
    taskDetailLink (id) {
      return "/tasks/" + id
    }
  },
  created() {
    axios
      .get(process.env.VUE_APP_CAMUNDA_PROXY_URL + '/tasks')
      .then(resp => {
        this.tasks = resp.data
      })
  }
}
</script>
