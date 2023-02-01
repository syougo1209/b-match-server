import {NavigationContainer } from '@react-navigation/native'
import {createNativeStackNavigator} from '@react-navigation/native-stack'
import { ConversationsScreen } from './screens/ConversationsScreen'
import { ProfileScreen } from './screens/ProfileScreen'
import { Screens } from './Screens'
import { RegistrationScreen } from './screens/RegistrationScreen'
import { NativeBaseProvider} from "native-base";
import { theme } from './styles/theme'

const Stack = createNativeStackNavigator<Screens>();

export default function App() {
  return (
    <NativeBaseProvider theme={theme}>
      <NavigationContainer
      >
        <Stack.Navigator
          screenOptions={{ headerShown: false }}
        >
          <Stack.Screen
            name="Conversations"
            component={ConversationsScreen}
            options={{title: 'Welcome'}}
          />
          <Stack.Screen name="Profile" component={ProfileScreen} />
          <Stack.Screen name="Registration" component={RegistrationScreen} />
        </Stack.Navigator>
      </NavigationContainer>
    </NativeBaseProvider>
  );
}
