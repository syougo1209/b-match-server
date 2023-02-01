import { createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import { ProfileScreen } from '../../../screens/ProfileScreen'
import { ConversationsScreen } from '../../../screens/ConversationsScreen'

const Tab = createBottomTabNavigator();

export const BottomMenu =()=> {
  return (
    <Tab.Navigator
    >
      <Tab.Screen
        name="ProfileScreen"
        component={ProfileScreen}
      />
      <Tab.Screen
        name="ConversationsScreen"
        component={ConversationsScreen}
      />
    </Tab.Navigator>
  );
}
