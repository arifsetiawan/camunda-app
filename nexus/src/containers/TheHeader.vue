<template>
  <CHeader fixed with-subheader light>
    <CToggler
      in-header
      class="ml-3 d-lg-none"
      @click="$store.commit('toggleSidebarMobile')"
    />
    <CToggler
      in-header
      class="ml-3 d-md-down-none"
      @click="$store.commit('toggleSidebarDesktop')"
    />
    <CHeaderBrand class="mx-auto d-lg-none" to="/">
      <CImg src="img/nexus.png" width="140" />
    </CHeaderBrand>
    <CHeaderNav class="d-md-down-none mr-auto">
    </CHeaderNav>
    <CHeaderNav class="mr-4">
      <CHeaderNavItem class="px-3">
        <CHeaderNavLink>
          {{userName}}
        </CHeaderNavLink>
      </CHeaderNavItem>
      <TheHeaderDropdownAccnt/>
    </CHeaderNav>
    <CSubheader class="px-3">
      <CBreadcrumbRouter class="border-0 mb-0"/>
    </CSubheader>
  </CHeader>
</template>

<script>
import TheHeaderDropdownAccnt from './TheHeaderDropdownAccnt'

export default {
  name: 'TheHeader',
  components: {
    TheHeaderDropdownAccnt
  },
  computed: {
    userName () {
      return this.$store.state.user.profile.firstName + " " + this.$store.state.user.profile.lastName
    }
  },
  created() {
    if (localStorage.getItem('user-profile')) {
      this.$store.state.user.profile = JSON.parse(localStorage.getItem('user-profile'))
    }
  }
}
</script>
