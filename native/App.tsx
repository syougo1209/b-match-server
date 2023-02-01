import { Provider as PaperProvider} from 'react-native-paper'
import {NavigationContainer } from '@react-navigation/native'
import {createNativeStackNavigator} from '@react-navigation/native-stack'
import { ConversationsScreen } from './screens/ConversationsScreen'
import { ProfileScreen } from './screens/ProfileScreen'
import { Screens } from './Screens'

const Stack = createNativeStackNavigator<Screens>();

export default function App() {
  return (
    <PaperProvider>
      <NavigationContainer
      >
        <Stack.Navigator
          screenOptions={{ headerShown: false }}
          initialRouteName="Conversations"
        >
          <Stack.Screen
            name="Conversations"
            component={ConversationsScreen}
            options={{title: 'Welcome'}}
          />
          <Stack.Screen name="Profile" component={ProfileScreen} />
        </Stack.Navigator>
      </NavigationContainer>
    </PaperProvider>
  );
}
