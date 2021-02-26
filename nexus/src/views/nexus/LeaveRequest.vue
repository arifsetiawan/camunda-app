<template>
  <div>
    <CRow>
      <CCol lg="12">
        <CCard>
            <CCardHeader>
                <strong>Company </strong><small>Form</small>
            </CCardHeader>
          <CCardBody>
            <CInput
              label="Company"
              placeholder="Enter your company name"
            />
            <CInput
              label="VAT"
              placeholder="PL1234567890"
            />
            <CInput
              label="Street"
              placeholder="Enter street name"
            />
            <CRow>
              <CCol sm="8">
                <CInput
                  label="City"
                  placeholder="Enter your city"
                />
              </CCol>
              <CCol sm="4">
                <CInput
                  label="Postal code"
                  placeholder="Postal code"
                />
              </CCol>
            </CRow>
            <CInput
              label="Country"
              placeholder="Country name"
            />
          </CCardBody>
        </CCard>
      </CCol>
    </CRow>
  </div>
</template>

<script>
export default {
  name: 'LeaveRequest'
}
</script>
