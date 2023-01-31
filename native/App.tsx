import { StatusBar } from 'expo-status-bar';
import { Text, View } from 'react-native';
import { Provider as PaperProvider, Appbar} from 'react-native-paper'

export default function App() {
  return (
    <PaperProvider>
      <Appbar.Header>
    <Appbar.BackAction onPress={() => {}} />
    <Appbar.Content title="Title" />
    <Appbar.Action icon="calendar" onPress={() => {}} />
    <Appbar.Action icon="magnify" onPress={() => {}} />
  </Appbar.Header>
      <View>
        <Text>Open up Apx to start working on your app!</Text>
        <StatusBar style="auto" />
      </View>
    </PaperProvider>
  );
}
