// Code generated by go-bluetooth generator DO NOT EDIT.

/*
BlueZ D-Bus Admin Policy API description [admin-policy-api.txt]
This API provides methods to control the behavior of bluez as an administrator.

Interface AdminPolicySet1 provides methods to set policies. Once the policy is
set successfully, it will affect all clients and stay persistently even after
restarting Bluetooth Daemon. The only way to clear it is to overwrite the
policy with the same method.

Interface AdminPolicyStatus1 provides readonly properties to indicate the
current values of admin policy.



*/
package admin_policy
