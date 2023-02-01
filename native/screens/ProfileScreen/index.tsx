import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { SafeAreaContainer } from '../../components/elements/SafeAreaContainer';

export const ProfileScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()
  return (
    <SafeAreaContainer>
      <Button
        title="Go n"
        onPress={() =>{
          console.log("hoge")
          navigation.navigate('Conversations')
        }}
      />
    </SafeAreaContainer>
  );
};
