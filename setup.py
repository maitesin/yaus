from distutils.core import setup

setup(name='short',
      version='0.1',
      description='URL shortener service',
      author='Oscar Forner Martinez',
      author_email='oscar.forner.martinez@gmail.com',
      url='https://www.gitlab.com/maitesin/short',
      py_modules=['short'],
      extras_require={"test": ["pytest", "coverage"]},
     )
