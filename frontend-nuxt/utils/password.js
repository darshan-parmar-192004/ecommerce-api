export function usePasswordStrength(passwordRef) {
  const passwordStrength = ref({
    level: 1,
    text: 'Weak',
    color: '#ff5252'
  });

  watch(
    passwordRef,
    (newPassword) => {
      let score = 0;
      const pwd = newPassword || '';
      
      if (pwd.length >= 6) score++;
      if (pwd.length >= 8) score++;
      if (/[a-z]/.test(pwd) && /[A-Z]/.test(pwd)) score++;
      if (/\d/.test(pwd)) score++;
      if (/[^a-zA-Z0-9]/.test(pwd)) score++;

      passwordStrength.value = {
        level: Math.min(3, score),
        text: score <= 2 ? 'Weak' : 'Strong',
        color: score === 3 ? '#4caf50' : '#ff5252'
      };
    }
  );

  return passwordStrength;
}