import { useRegisterProfileContext } from '../../../context/useRegisterProfileContext';
import { useController, useFormContext } from 'react-hook-form';
import { NickNameForm } from '../../../components/NickNameForm';

export const NickNameFormContainer = () => {
  const {onPressBackButton, onPressSubmitButton} = useRegisterProfileContext()
  const { control } = useFormContext()
  const { field } = useController({name: 'nickname', control})

  return (
     <NickNameForm
        onPressBackButton={onPressBackButton}
        onPressSubmitButton={onPressSubmitButton}
        onChangeText={field.onChange}
        onBlur={field.onBlur}
        inputRef={field.ref}
        fieldValue={field.value}
      />
  )
}
