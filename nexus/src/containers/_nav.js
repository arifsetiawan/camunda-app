export default [
  {
    _name: 'CSidebarNav',
    _children: [
      {
        _name: 'CSidebarNavItem',
        name: 'Home',
        to: '/home',
        icon: 'cil-home',
      },
      {
        _name: 'CSidebarNavTitle',
        _children: ['Tasks']
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Tasks',
        to: '/tasks',
        icon: 'cil-list-rich'
      },
      {
        _name: 'CSidebarNavTitle',
        _children: ['Forms']
      },
      {
        _name: 'CSidebarNavItem',
        name: 'Leave Request',
        to: '/workflow/leave',
        icon: 'cil-pencil'
      }
    ]
  }
]