import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { SafeAreaContainer } from '../../components/elements/SafeAreaContainer';

export const ProfileScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Registration'>>()
  return (
    <SafeAreaContainer>
      <Button
        title="Go Register"
        onPress={() =>{
          console.log("hoge")
          navigation.navigate('Registration')
        }}
      />
    </SafeAreaContainer>
  );
};
