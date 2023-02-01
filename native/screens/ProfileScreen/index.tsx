import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { Container, Box } from 'native-base'

export const ProfileScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Registration'>>()
  return (
    <Container safeArea>
      <Button
        title="Go Register"
        onPress={() =>{
          console.log("hoge")
          navigation.navigate('Registration')
        }}
      />
      <Box>Hello</Box>
    </Container>
  );
};
