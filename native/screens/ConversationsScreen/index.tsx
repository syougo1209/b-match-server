import { Button } from 'react-native'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { Box, Container } from "native-base"

export const ConversationsScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Profile'>>()
  return (
    <Container safeArea>
      <Button
        title="Go ton"
        onPress={() =>
          navigation.navigate('Profile')
        }
      />
    </Container>
  );
};
